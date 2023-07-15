# concurrency - mutex

Channel communication is great for synchronizing routine execution. But, if you don't need execution synchronization and only need data synchronization, Go offers a mutex. The example shows this by starting 1000 routines all interested in the same data.
