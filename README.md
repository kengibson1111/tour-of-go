# tour-of-go ![tour-of-go](/images/icons8-golang-240.png)

This is a companion repo for Go's [Tour of Go](https://go.dev/tour/welcome/1).[^1] Definitely use the code window on the tour. This repo contains mostly code from the tour, and at times I add more code based on ideas from the lessons. Code from the welcome pages is not included.

## Lessons

* goroutines - introduction to Go routines. Each routine is a lightweight thread accessing resources potentially
  shared by other routines.

* channels - these are like Java streams. You send and receive values with the operator, <-. You have to use
  make() like with maps and slices. Sends and receives block until the other side is ready, and that makes it very
  useful for routine synchronization without explicit locks. So the sample distributes work among 2 routines.
  When each routine is done, it sends an int on a channel. Both x and y ints are waiting to receive from the
  channel. 2 senders, 2 receivers - remember the blocking. So I start 2 threads to do work and then I wait in the
  caller until both threads are done. AND I have to receive results from both threads. Do the same in Java
  and compare the code :).

* channels (buf) - you can pass in a 2nd argument to make() which specifies a buffer length. When you do this,
  golang uses a buffer with the channel. Sends block until the buffer is full. Receives block if the buffer
  is empty. So try a couple of code mods and see what happens. First, try sending a 3rd int on the channel before
  the first Println(). Back to the original code (without the 3rd int send), try adding a 3rd Println(). The
  advantage of a buffer is to allow heavy traffic even though a receiver might not be ready. Without a buffer,
  the channel blocks on just a single send. With a buffer, sends can pile up in a buffer and receivers can pull
  when they are ready.

* channels (close) - back to the keyword range used in a loop. It works on a channel, but the channel has to
  be closed for the loop to end. Channels don't have to be closed unless a receiver needs to be told. And
  it is a good practice to only have a sender close a channel. Sending on a closed channel causes a runtime
  panic. The sample creates a buffered channel (and the cap() function can be used on a buffered channel).
  A routine is created and then main() waits on the range keyword. When the buffer has 10 ints, the loop
  activates because range can receive. Just before the routine shuts down, it closes the channel. This tells
  range to kick out of the loop. Cool. Again, write the same functionality in Java and compare the code :).

* channels (select) - another way to do the previous lesson without using close(). select waits until a
  channel operation is ready, and if multiple channel operations are ready the channel operation order
  is not guaranteed. Create 2 channels - not buffered. One channel is the main processing channel and the other
  is a "quit" channel. fibonacci() has an endless loop sending ints on the main processing channel
  and receiving ints on the "quit" channel. When one "quit" int is received, fibonacci() ends which also
  ends main(). The Go routine receives 10 ints from the main processing channel and sends one int on the
  "quit" channel.

* channels (selectdef) - just showing that a default can be used in a select. This will activate when all
  channel operations on the select are blocked.

* EXERCISE (concur) - main exercise for the Concurrency section of the tour. Maybe not the best code, but
  it was easy to implement and satisfies requirements.

* mutex - so channel communication is great for synchronizing routine execution. But if you don't need
  execution synchronization and only need data synchronization, golang offers a mutex. The sample
  shows this by starting 1000 routines all interested in the same data.

* EXERCISE (mutex) - last exercise building on previous lesson.

## keys-remote.sh

In a DEV container separate from the container used by VS Code, use [keys-remote.sh](https://github.com/kengibson1111/tour-of-go/blob/master/keys-remote.sh) to validate remote Git connectivity. If results look good, then any `git remote -v` error from VS Code is not related to how Git is interacting with Windows 11 Pro.

## Summary

The tour is pretty fun. Hopefully it will help others get off the ground and start exploring Go.

***

[^1]: Icon provided by [Icons8](https://icons8.com/).
