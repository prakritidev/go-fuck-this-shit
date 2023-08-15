# Go Getting started notes


# Interface 


# Refelection 

# What is sturct ?

# What the fucks is concurrency and parallel ?

concurrency is a software development construct for better use of CPU time. 

concurrency is composition of independent execting tasks. But not neccessarly consuming CPU time. 

# what is go ?

Go is a concurrent language. Their idea is to use pass the data though channels rather then sharing memory between concurrent process and applying logics on the memory. 

It is desinged for developing system software 
# Go Channels 

They are first class values, very similar to what messaging queues are in system desing. 
A channel provied a way to communicate between two Goroutines.

We can can millions of channels as the are not exaclty threads, this is very similar to what java has as virtual threads(Not exaclty but very similat concept wise).

Sending and recieveing is a blocking operation, and a Synchronization operation in a single operation.


# Concurrent programming patterns. 

Channels are first class values. 

#### Generator Functions
It returns channel. if we wait for our driver method the channel is blocked due to the nature of channels. For mitigaing that we have Multiplexer functions. 


#### Multiplexer Functions


This is nothing special, just create a new fucntion that takes a list if channels and returns a single channel and your driver function can listen to that channel.

This will break your sequencing like aws standard queues does. 
#### Restoring sequencing. 


# Select statement combination

1. Fan-in 
2. Timeout using Select
3. Quit channel




#### daisy chain, 




# Goroutines

create individial task that will run independent of you main method, howver you main method has to wait forroutines to get completed otherwise it will end wihtout processing anything. 

This is not a classic java joinAndFork implementation.


### Advanced Go concurrency patterns. 



Even though we don't have mutex and semaphore, go still has deadlocks. 


