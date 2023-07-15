# concurrency - channels (close)

Back to the keyword range used in a `for` loop. It works on a channel, but the channel has to be closed to end the loop. Channels don't have to be closed unless a receiver needs to be told. And it is a best practice to only have a sender close a channel. Sending on a closed channel causes a runtime panic.

The example creates a buffered channel (and the cap() function can be used on a buffered channel). A goroutine is created and then main() waits on the range keyword. When the buffer has 10 ints, the loop activates because range can receive. Just before the routine shuts down, it closes the channel. This tells range to kick out of the loop. I don't think the code would be as clean in Java.
