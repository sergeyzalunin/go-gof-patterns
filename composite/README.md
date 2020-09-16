## **Composition**
A mechanism for threating individual (scalar objects) and compositions of objects in a uniform manner

## Motivation to use composition

* Objects use other objects' fields/methods through embedding
* Composition lets us make compound objects
  * E.g., a mathematica expression composed of simple expressions; or
  * A shape group made of several different shapes
* Composite design pattern is used to threat both single (scalar) and composite objects uniformly
* I.e., Foo and []Foo have common APIs

## Summary

* Objects can use other objects via composition
* Some composed and singular objects need similar/identical behaviors
* Compose design pattern lets us treat both types of objects uniformly
* Iteration supported with the Iterator design pattern