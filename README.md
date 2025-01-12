## Interface

The interface we want to define:


* EDBT -> root component, setup a behaviour tree
* Nodes -> each behaviour node.
    * Composite nodes
        * Selector (OR behaviour)
        * Sequencer (AND behaviour)
        * Parallel (concurrent behaviour)
    * Decorators (decorate tasks)
        * BOD - blackboard observer decorator
        * Conditional?
        * Inversion?
    * Coordination
        * Request Handler
        * Soft Message Sender
        * Hard Message Sender
    * Event-driven
        * ?
    * Tasks
        * Abstract behaviours
