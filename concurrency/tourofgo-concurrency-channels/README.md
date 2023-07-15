# concurrency - channels

These are similar to Java streams. You send and receive values with the operator, <-. You have to use make() like with maps and slices. Sends and receives block until the other side is ready, and that makes it very useful for routine synchronization without explicit locks.
  
The example distributes work among 2 routines. When each routine is done, it sends an int on a channel. Both x and y ints are waiting to receive from the channel. 2 senders, 2 receivers - remember the blocking. So 2 goroutines are started to do work and then the main code blocks until both goroutines are done. AND I have to receive results from both threads. I don't know if that code is as clean in Java.
