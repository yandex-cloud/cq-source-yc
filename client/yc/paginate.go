package yc

import (
	"context"

	"google.golang.org/grpc"
)

// Paginate drives a standard YC List RPC (page_token / next_page_token) to
// completion, streaming every item to res. It is the concise replacement for the
// hand-written fetch loop of SDK-less services — the generated gRPC client already
// gives us List, GetNextPageToken and the repeated-field getter, so a resolver is:
//
//	cc, err := c.GRPCConn(ctx, "yandex.cloud.ai.files.v1.FileService.List")
//	if err != nil {
//		return err
//	}
//	cl := files.NewFileServiceClient(cc)
//	return yc.Paginate(ctx, res,
//		func(t string) *files.ListFilesRequest { return &files.ListFilesRequest{FolderId: c.FolderId, PageToken: t} },
//		cl.List,
//		(*files.ListFilesResponse).GetFiles,
//		(*files.ListFilesResponse).GetNextPageToken,
//	)
//
// Type parameters are inferred from the arguments. newReq builds a fresh request
// carrying the supplied page token; list is the client's List method; items and
// nextToken are the response getters (method expressions).
func Paginate[Req, Resp, Item any](
	ctx context.Context,
	res chan<- any,
	newReq func(pageToken string) Req,
	list func(context.Context, Req, ...grpc.CallOption) (Resp, error),
	items func(Resp) []Item,
	nextToken func(Resp) string,
) error {
	for token := ""; ; {
		resp, err := list(ctx, newReq(token))
		if err != nil {
			return err
		}
		for _, item := range items(resp) {
			res <- item
		}
		if token = nextToken(resp); token == "" {
			return nil
		}
	}
}
