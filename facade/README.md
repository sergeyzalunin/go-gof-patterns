## **Facade**
Provides a simple, easy to understand/user interface over a large and sophisticated body of code


## Motivation to use composition

* Balancing complexity and presentation/usability
* E.g.: Typical home:
  * Many subsystems (electrical sanitation)
  * Complex internal structure (e.g., floor layers)
  * End user is not exposed to internals
* Same with software!
  * Many systems working to provide flexibility, but...
  * API consumers want it it "just work"

## Summary

* Build a Facade to provide a simplified API over a set of components
* May wish to (optionally) expose internals through the facade
* May allow users to `escalate` to use more complex APIs if they need to