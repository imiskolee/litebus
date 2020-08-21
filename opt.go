package litebus

type PublishOption struct {
	FailFirst bool //fail the message when error from handler coming and stop run next handler.
	Parallel  bool
}
