//class Circle
public class Circle extends Shape{
	public Circle(double radius){
		setRadius(radius);
	}
	
	public Circle(){
		//do nothing
	}
	
	//calculate area of circle
	@Override
	public double area(){
		return Math.PI * Math.pow(getRadius(), 2);
	}
	
	//calculate circumference of circle
	@Override
	public double perimeter(){
		return 2 * Math.PI * getRadius();
	}
	

	//return String representation of area
	@Override
	public String toString() {
		return String.format("Circumference : %.2f\nArea : %.2f\n", perimeter(), area());
	}
}

