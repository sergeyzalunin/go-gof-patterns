## **Prototype**
When piecewise object construction is complicated, provide an API for doing it succinctly

## Motivation to use prototype

* Some objects are simple and can be created in a single constructor call
* Other objects require a lot of ceremony to create
* Having a factory function with 10 arguments is not productive
* Builder provides an API for constructing an object step-by-step

## Summary

* A builder is a separate component used for building an object
* To make builder fluent, return the receiver - allows chaining
* Different facets of an object can be buit with different builders working in tandem via a common struct