//class Square
public class Square extends Shape{

	public Square(double l1, double l2){
		setLengthAndWidth(l1, l2);
	}
	
	public Square(){
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

	//return string representation of perimeter and area
	@Override
	public String toString(){
		return String.format("Perimeter : %.2f\nArea : %.2f\n",perimeter(),area()); 
	}
}
