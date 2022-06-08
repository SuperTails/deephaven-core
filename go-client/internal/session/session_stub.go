package session

import (
	"context"
	"fmt"
	"sync"
	"time"

	sessionpb2 "github.com/deephaven/deephaven-core/go-client/internal/proto/session"

	"google.golang.org/grpc/metadata"
)

type tokenResp struct {
	Token []byte
	Error error
}

// Stores the current session token and sends periodic keepalive messages
type refresher struct {
	ctx         context.Context
	sessionStub sessionpb2.SessionServiceClient

	tokenMutex *sync.Mutex
	token      *tokenResp

	cancelCh chan struct{}

	timeoutMillis int64
}

func (ref *refresher) refreshLoop() {
	for {
		select {
		case <-ref.cancelCh:
			return
		case <-time.After(time.Duration(ref.timeoutMillis) * time.Millisecond):
			err := ref.refresh()
			if err != nil {
				return
			}
		}
	}
}

func startRefresher(ctx context.Context, sessionStub sessionpb2.SessionServiceClient, tokenMutex *sync.Mutex, token *tokenResp, cancelCh chan struct{}) error {
	handshakeReq := &sessionpb2.HandshakeRequest{AuthProtocol: 1, Payload: [](byte)("hello godeephaven")}
	handshakeResp, err := sessionStub.NewSession(ctx, handshakeReq)
	if err != nil {
		return err
	}

	ctx = metadata.NewOutgoingContext(context.Background(), metadata.Pairs("deephaven_session_id", string(handshakeResp.SessionToken)))

	{
		tokenMutex.Lock()
		token.Token = handshakeResp.SessionToken
		tokenMutex.Unlock()
	}

	timeoutMillis := handshakeResp.TokenExpirationDelayMillis / 2

	ref := refresher{
		ctx:         ctx,
		sessionStub: sessionStub,

		tokenMutex: tokenMutex,
		token:      token,

		cancelCh: cancelCh,

		timeoutMillis: timeoutMillis,
	}

	go ref.refreshLoop()

	return nil
}

func (ref *refresher) refresh() error {
	ref.tokenMutex.Lock()
	oldToken := make([]byte, len(ref.token.Token))
	copy(oldToken, ref.token.Token)
	ref.tokenMutex.Unlock()

	handshakeReq := &sessionpb2.HandshakeRequest{AuthProtocol: 0, Payload: oldToken}
	handshakeResp, err := ref.sessionStub.RefreshSessionToken(ref.ctx, handshakeReq)

	if err != nil {
		ref.tokenMutex.Lock()
		ref.token.Error = err
		ref.tokenMutex.Unlock()
		fmt.Println("Failed to refresh token:", err)
		return err
	} else {
		ref.tokenMutex.Lock()
		ref.token.Token = handshakeResp.SessionToken
		ref.tokenMutex.Unlock()
	}

	ref.timeoutMillis = handshakeResp.TokenExpirationDelayMillis / 2

	fmt.Println("New timeout millis:", ref.timeoutMillis)

	return nil
}

type SessionStub struct {
	session *Session

	stub sessionpb2.SessionServiceClient

	tokenMutex *sync.Mutex
	token      *tokenResp

	cancelCh chan struct{}
}

// Performs the first handshake to get a session token.
//
func NewSessionStub(ctx context.Context, session *Session) (SessionStub, error) {
	stub := sessionpb2.NewSessionServiceClient(session.GrpcChannel())

	cancelCh := make(chan struct{})

	tokenMutex := &sync.Mutex{}
	tokenResp := &tokenResp{}

	err := startRefresher(ctx, stub, tokenMutex, tokenResp, cancelCh)
	if err != nil {
		return SessionStub{}, err
	}

	hs := SessionStub{
		tokenMutex: tokenMutex,
		token:      tokenResp,

		cancelCh: cancelCh,
	}

	return hs, nil
}

func (hs *SessionStub) Token() []byte {
	hs.tokenMutex.Lock()
	if hs.token.Error != nil {
		panic("TODO: Error in refreshing token")
	}
	token := make([]byte, len(hs.token.Token))
	copy(token, hs.token.Token)
	hs.tokenMutex.Unlock()

	return token
}

func (hs *SessionStub) Close() {
	if hs.cancelCh != nil {
		close(hs.cancelCh)
		hs.cancelCh = nil
	}
}
