# concurrency - channels (buffered)

A 2nd argument to make() can be passed which specifies a buffer length. When a buffer length is used, Go uses a buffer with the channel. Sends block until the buffer is full. Receives block if the buffer is empty. Try a couple of code mods and see what happens. First, try sending a 3rd int on the channel before the first Println(). Back to the original code (without the 3rd int send), try adding a 3rd Println().

The advantage of a buffer is to allow heavy traffic even though a receiver might not be ready. Without a buffer, the channel blocks on just a single send. With a buffer, sends can pile up in a buffer and receivers can pull when they are ready.
