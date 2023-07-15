# tour-of-go ![tour-of-go](/images/icons8-golang-240.png)

This is a companion repo for Go's [Tour of Go](https://go.dev/tour/welcome/1).[^1] Definitely use the code window on the tour. This repo contains mostly code from the tour, and at times I add more code based on ideas from the lessons. Code from the welcome pages is not included.

## Lessons

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
