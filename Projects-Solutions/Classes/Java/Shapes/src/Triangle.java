//class Triangle
public class Triangle extends Shape{
	public Triangle(double side1, double side2, double side3){
		setSides(side1, side2, side3);
	}
	
	public Triangle(){
		//do nothing
	}
	
	/*calculate area using formula 
	 * A = square root of (p(p-a)(p-b)(p-c))
	 * where p = (1/2)(a + b + c) and a,b,c are sides of the triangle
	 * */
	@Override
	public double area(){
		double p = (1/2.0) * (getSide1() + getSide2() + getSide3());
		return Math.sqrt(p * (p-getSide1()) * (p-getSide2()) * (p-getSide3()));
	}
	
	//calculate perimeter
	@Override
	public double perimeter(){
		return getSide1() + getSide2() + getSide3();
	}
	
	//return String representation of area and perimeter 
	@Override
	public String toString(){
		return String.format("Perimeter : %.2f\nArea : %.2f", perimeter(), area());
	}
}
