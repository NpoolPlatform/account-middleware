package deposit

import (
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/deposit"
)

func trace(span trace1.Span, in *npool.AccountReq, index int) trace1.Span {
	return span
}

func Trace(span trace1.Span, in *npool.AccountReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	return span
}

func TraceMany(span trace1.Span, infos []*npool.AccountReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
