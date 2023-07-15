# concurrency - channels (select)

This is a way to coordinate multiple-channel processing in a single function using a `select` statement. `select` waits until a channel operation is ready, and if multiple channel operations are ready the channel operation order is not guaranteed.

What I like about this example is that one channel is for outgoing data from fibonacci() and another channel is incoming for the quit message. Both channels are not buffered, so as soon as a single value is on the outgoing channel from fibonacci(), processing occurs in the inline goroutine in main.
