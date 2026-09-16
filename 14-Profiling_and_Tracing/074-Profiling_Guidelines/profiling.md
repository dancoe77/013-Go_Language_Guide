## Profiling Code

We can use the go tooling to inspect and profile our programs. Profiling is more of a journey and detective work. It requires some understanding about the application and expectations. The profiling data in and of itself is just raw numbers. We have to give it meaning and understanding.

## The Basics of Profiling

_"Those who can make you believe absurdities can make you commit atrocities" - Voltaire_

### How does a profiler work?

A profiler runs your program and configures the operating system to interrupt it at regular intervals. This is done by sending SIGPROF to the program being profiled, which suspends and transfers execution to the profiler. The profiler then grabs the program counter for each executing thread and restarts the program.

### Profiling do's and don't's

Before you profile, you must have a stable environment to get repeatable results.

* The machine must be idle—don't profile on shared hardware, don't browse the web while waiting for a long benchmark to run.
* Watch out for power saving and thermal scaling.
* Avoid virtual machines and shared cloud hosting; they are too noisy for consistent measurements.

If you can afford it, buy dedicated performance test hardware. Rack them, disable all the power management and thermal scaling and never update the software on those machines.

For everyone else, have a before and after sample and run them multiple times to get consistent results.

### Types of Profiling

**CPU profiling**  
CPU profiling is the most common type of profile. When CPU profiling is enabled, the runtime will interrupt itself every 10ms and record the stack trace of the currently running goroutines. Once the profile is saved to disk, we can analyse it to determine the hottest code paths. The more times a function appears in the profile, the more time that code path is taking as a percentage of the total runtime.

**Memory profiling**  
Memory profiling records the stack trace when a heap allocation is made. Memory profiling, like CPU profiling is sample based. By default memory profiling samples 1 in every 1000 allocations. This rate can be changed. Stack allocations are assumed to be free and are not tracked in the memory profile. Because of memory profiling is sample based and because it tracks allocations not use, using memory profiling to determine your application's overall memory usage is difficult.

**Block profiling**  
Block profiling is quite unique. A block profile is similar to a CPU profile, but it records the amount of time a goroutine spent waiting for a shared resource. This can be useful for determining concurrency bottlenecks in your application. Block profiling can show you when a large number of goroutines could make progress, but were blocked.

Blocking includes:

* Sending or receiving on an unbuffered channel.
* Sending to a full channel, receiving from an empty one.
* Trying to Lock a sync.Mutex that is locked by another goroutine.
* Block profiling is a very specialised tool, it should not be used until you believe you have eliminated all your CPU and memory usage bottlenecks.

**One profile at at time**  
Profiling is not free. Profiling has a moderate, but measurable impact on program performance—especially if you increase the memory profile sample rate. Most tools will not stop you from enabling multiple profiles at once. If you enable multiple profiles at the same time, they will observe their own interactions and skew your results.

**Do not enable more than one kind of profile at a time.**

### Hints to interpret what you see in the profile

If you see lots of time spent in `runtime.mallocgc` function, the program potentially makes excessive amount of small memory allocations. The profile will tell you where the allocations are coming from. See the memory profiler section for suggestions on how to optimize this case.

If lots of time is spent in channel operations, `sync.Mutex` code and other synchronization primitives or System component, the program probably suffers from contention. Consider to restructure program to eliminate frequently accessed shared resources. Common techniques for this include sharding/partitioning, local buffering/batching and copy-on-write technique.

If lots of time is spent in `syscall.Read/Write`, the program potentially makes excessive amount of small reads and writes. Bufio wrappers around os.File or net.Conn can help in this case.

If lots of time is spent in GC component, the program either allocates too many transient objects or heap size is very small so garbage collections happen too frequently.

* Large objects affect memory consumption and GC time, while large number of tiny allocations affects execution speed.

* Combine values into larger values. This will reduce number of memory allocations (faster) and also reduce pressure on garbage collector (faster garbage collections).

* Values that do not contain any pointers are not scanned by garbage collector. Removing pointers from actively used value can positively impact garbage collection time.

## Rules of Performance

1) Never guess about performance.
2) Measurements must be relevant.
3) Profile before you decide something is performance critical.
4) Test to know you are correct.

## Installing Tools

**hey**  
hey is a modern HTTP benchmarking tool capable of generating the load you need to run tests. It's built using the Go language and leverages goroutines for behind the scenes async IO and concurrency.

	go get -u github.com/rakyll/hey