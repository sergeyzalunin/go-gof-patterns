## **Prototype**
A partially or fully initialized object that you copy (clone) and make use of

## Motivation to use prototype

* Complicated objects (e.g., cars) aren't designed from scratch
    * The reiterate existing designs
* An existing (partially or fully constructed) design is a Prototype
* We make a copy of the prototype and customize it
    * Requires 'deep copy' support
* We make the cloning convenient (e.g via a Factory)

## Summary

* To implement a prototype, partially construct an object and store it somewhere
* Deep copy the prototype
* Customize the resulting instance
* A prototype factory provides a convenient API for using prototypes