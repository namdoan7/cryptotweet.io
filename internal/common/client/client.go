package client

import (
	"crypto/tls"
	"crypto/x509"
	"os"
	"strconv"

	profilepb "github.com/levinhne/cryptotweet.io/internal/common/genproto/profile"
	tagpb "github.com/levinhne/cryptotweet.io/internal/common/genproto/tag"
	tweetpb "github.com/levinhne/cryptotweet.io/internal/common/genproto/tweet"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func NewTweetClient() (client tweetpb.TweetServiceClient, close func() error, err error) {
	grpcAddr := os.Getenv("TWEET_GRPC_ADDR")
	if grpcAddr == "" {
		return nil, func() error { return nil }, errors.New("empty env TWEET_GRPC_ADDR")
	}

	opts, err := grpcDialOpts(grpcAddr)

	if err != nil {
		return nil, func() error { return nil }, err
	}

	conn, err := grpc.Dial(grpcAddr, opts...)
	if err != nil {
		return nil, func() error { return nil }, err
	}

	return tweetpb.NewTweetServiceClient(conn), conn.Close, nil
}

func NewProfileClient() (client profilepb.ProfileServiceClient, close func() error, err error) {
	grpcAddr := os.Getenv("PROFILE_GRPC_ADDR")
	if grpcAddr == "" {
		return nil, func() error { return nil }, errors.New("empty env PROFILE_GRPC_ADDR")
	}

	opts, err := grpcDialOpts(grpcAddr)

	if err != nil {
		return nil, func() error { return nil }, err
	}

	conn, err := grpc.Dial(grpcAddr, opts...)
	if err != nil {
		return nil, func() error { return nil }, err
	}

	return profilepb.NewProfileServiceClient(conn), conn.Close, nil
}

func NewTagClient() (client tagpb.TagServiceClient, close func() error, err error) {
	grpcAddr := os.Getenv("TAG_GRPC_ADDR")
	if grpcAddr == "" {
		return nil, func() error { return nil }, errors.New("empty env TAG_GRPC_ADDR")
	}

	opts, err := grpcDialOpts(grpcAddr)

	if err != nil {
		return nil, func() error { return nil }, err
	}

	conn, err := grpc.Dial(grpcAddr, opts...)
	if err != nil {
		return nil, func() error { return nil }, err
	}

	return tagpb.NewTagServiceClient(conn), conn.Close, nil
}

func grpcDialOpts(grpcAddr string) ([]grpc.DialOption, error) {
	return []grpc.DialOption{grpc.WithInsecure()}, nil
	if noTLS, _ := strconv.ParseBool(os.Getenv("GRPC_NO_TLS")); noTLS {
		return []grpc.DialOption{grpc.WithInsecure()}, nil
	}

	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		return nil, errors.Wrap(err, "cannot load root CA cert")
	}
	creds := credentials.NewTLS(&tls.Config{
		RootCAs:    systemRoots,
		MinVersion: tls.VersionTLS12,
	})

	return []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
		// grpc.WithPerRPCCredentials(newMetadataServerToken(grpcAddr)),
	}, nil
}
