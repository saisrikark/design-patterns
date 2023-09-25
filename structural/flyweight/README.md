# flyweight

<p>
Provides a way to decrease object count. <br/>
Thereby decreases the amount of RAM used. <br/>
Generally used when large number of objects are created. <br/>
Flyweights are immutable, they cannot be modified once constructed. <br/>
<br/>
For flyweight properties can be divided to intrinsic and extrinsic. <br/>
Intrinsic - fields that can be stored inside the object, these generally don't change a lot and are few combinations in number, this is immutable. <br/>
Extrinsic - fields that are stored outside the object, that are mutable as well;
</p>
