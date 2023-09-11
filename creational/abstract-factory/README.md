# abstract factory

<p> 
An interface to create families of related/dependent objects without specifying their concrete classes. <br />
It's a super factory that creates other factories. <br />
You can be sure that the products you’re getting from a factory are compatible with each other. <br />
</p>

Elements
1. Abstract factory
2. Concrete factory
3. Abstract product #1
4. Abstract product #2
5. Abstract product ..
6. Abstract product #n
7. Concrete products for each abstract products

Example
A customer wants to pick a certain kind of home decor
They can choose Indian, American or Italian.
Each type of decor is a set composed of a chair, sofa and table.
Chairs, sofas and tables have their own behaviour.