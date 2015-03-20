import java.util.Scanner;

/* class Main : contains main method*/
public class Main {
	public static void main(String[] args){
		Scanner input = new Scanner(System.in);
		double radius, length, width;
		
		System.out.println("Select a shape");
		System.out.print("1.Square\n2.Rectangle\n3.Circle\n4.Triangle\n? ");
		int choice = Integer.parseInt(input.nextLine());
		
		switch(choice){
			case 1:
				System.out.println("Enter the length of one side of the square");
				length = Double.parseDouble(input.nextLine());
				Square square = new Square(length, length);
				System.out.println(square.toString());
				break;

			case 2:
				System.out.println("Enter the length and width of the rectangle");
				System.out.print("Length : ");
				length = Double.parseDouble(input.nextLine());
				System.out.print("Width : ");
				width = Double.parseDouble(input.nextLine());
				Rectangle rectangle = new Rectangle(length, width);
				System.out.println(rectangle.toString());
				break;

			case 3:
				System.out.println("Enter radius of the circle");
				radius = Double.parseDouble(input.nextLine());
				Circle circle = new Circle(radius);
				System.out.println(circle.toString());
				break;
				
			case 4:
				System.out.println("Enter the length of the sides of the triangle");
				System.out.print("side1: ");
				double s1 = Double.parseDouble(input.nextLine());
				System.out.print("side2: ");
				double s2 = Double.parseDouble(input.nextLine());
				System.out.print("side3: ");
				double s3 = Double.parseDouble(input.nextLine());
				Triangle triangle = new Triangle(s1, s2, s3);
				System.out.println(triangle.toString());
				break;
			default:
				System.out.println("Invalid choice");
		}
		input.close();

	}
}