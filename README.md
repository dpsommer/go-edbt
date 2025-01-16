## Interface

The interface we want to define (based on [Agis et al., 2020](https://www.sciencedirect.com/science/article/abs/pii/S0957417420302815) and [Champandard and Dunstan, 2012](https://www.gameaipro.com/GameAIPro/GameAIPro_Chapter06_The_Behavior_Tree_Starter_Kit.pdf)):


* EDBT -> root component, setup a behaviour tree
* Nodes -> each behaviour node.
    * Composite nodes
        * [x] Selector (OR behaviour)
        * [x] Sequencer (AND behaviour)
        * [x] Parallel (concurrent behaviour)
        * [ ] Filter/preconditions?
        * [ ] Monitor?
        * [ ] Active selector?
    * Decorators (decorate tasks)
        * [ ] BOD - blackboard observer decorator
        * [ ] Conditional?
        * [ ] Inversion?
    * Event-driven
        * [ ] ?
    * Coordination
        * [ ] Request Handler
        * [ ] Soft Message Sender
        * [ ] Hard Message Sender
    * Tasks
        * [ ] Abstract behaviours
