# Concurrency

### Checkpoint Synchronization [View](/concurrency/checkpoint-sync.go)
- This is a simple problem of synchronizing multiple tasks
- Consider a workshop when multiple workers assembling different parts of a mechanism
- When each of them completes its work, put the details together
- A worker who finishes its part first must wait for others before starting a new one
- Putting details together is the checkpoint at which tasks synchronize themselves before going them paths apart

### Dining Philosophers [View](/concurrency/dining-philosophers.go)
- This problem consists of five philosophers sit around a table eager to take a dinner
- Forks are placed between each pair of adjacent philosophers
- Each philosopher must alternately think and eat
- A philosopher can only eat when they have both forks at the same time (left and right)
- Each fork can be held only by one philosopher and so a philosopher can use the fork only if it's not being used by another philosopher
- After a philosopher has finished eating, they have to put down both forks in the table so they are available again
- A philosopher can take either the right or left available forks, but they have to wait to take both to start eating
- There is no restriction about the demand and supply
- The problem is to design a logic where no philosophers would starve

### Todo
- [ ] Producer Consumer Problem
- [ ] Sleeping Barber Problem
- [ ] Cigarette Smokers Problem
