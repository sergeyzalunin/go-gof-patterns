## **Builder**
A mechanism that decouples an interface (hierarchy) from an implementation (hierarchy)

## Motivation to use builder

* Bridge revents a `Cartesian product` complexety explosion
* Example:
    * Common type ThreadScheduler
    * Can be preemptive or cooperative
    * Can run on Windows or Unix
    * End up with a 2x2 scenario: WindosPTS, UnixPTS, WindowsCTS, UnixCTS
* Bridge pattern avoids the entity explosion

## Summary

* Decouple abstraction from implementation
* Both can exist as hierarchies
* A stronger form of encapsulation