//abstract class Shape
public abstract class Shape {
	private double length, width; //square and rectangle
	private double radius; //circle
	private double side1, side2, side3;	//triangle
	
	//set length and width
	public void setLengthAndWidth(double l, double w){
		length = l;
		width = w;
	}
	
	//return length
	public double getLength(){
		return length;
	}
	
	//return width
	public double getWidth(){
		return width;
	}
	
	//set radius
	public void setRadius(double r){
		radius = r;
	}

	//return radius
	public double getRadius(){
		return radius;
	}
	
	
	//set sides of triangle
	public void setSides(double s1, double s2, double s3){
		side1 = s1;
		side2 = s2;
		side3 = s3;
	}
	
	//return side1
	public double getSide1(){
		return side1;
	}
	
	//return side2
	public double getSide2(){
		return side2;
	}
	
	//return side3
	public double getSide3(){
		return side3;
	}
	
	//abstract methods ... to be overriden by concrete methods
	public abstract String toString();	//return String representation of the class
	public abstract double area();	//calculate area
	public abstract double perimeter();	//calculate perimeter
}
