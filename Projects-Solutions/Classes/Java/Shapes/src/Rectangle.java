//class Rectangle
public class Rectangle extends Shape{
	
	public Rectangle(double length, double width){
		setLengthAndWidth(length, width);
	}
	
	public Rectangle(){
		//do nothing
	}
	
	//calculate area
	@Override
	public double area(){
		return getLength() * getWidth();
	}
	
	//calculate perimeter
	@Override
	public double perimeter(){
		return 2 * (getLength() + getWidth());
	}

	//return String representation of area and perimeter
	@Override
	public String toString(){
		return String.format("Perimeter : %.2f\nArea : %.2f\n", perimeter(), area());
	}
}