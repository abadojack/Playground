/*contains Main method to demonstrate the working other classes*/
import java.util.Scanner;

public class Main {
    static Scanner input = new Scanner(System.in);
    public static void main(String[] args){
        //create different accounts and set random account numbers
        Account accChecking = new CheckingAccount(1234);
        Account accSaving = new SavingsAccount(1235);
        Account accBusiness = new BusinessAccount(1236);

        //set account balance for each of the accounts
        accBusiness.setAccountBalance(100000);
        accChecking.setAccountBalance(750000);
        accSaving.setAccountBalance(50000);


        //prompt user to select account type
        System.out.println("Select Account to use");
        System.out.println("1.Savings Account\n2.Checking Account\n3.Business Account\nEnter choice");
        int accountType = Integer.parseInt(input.nextLine());

        //perform operation on the account selected
        switch (accountType){
            case 1:
                accountOperations(accSaving);
                break;
            case 2:
                accountOperations(accChecking);
                break;
            case 3:
                accountOperations(accBusiness);
                break;
            default:
                System.out.println("Invalid choice");

        }
    }

    //operations the user can perform to their account
    public static void accountOperations(Account account){
        System.out.println("1.Check balance\n2.Withdrawal\n3.Deposit\n4.Exit\nEnter choice");
        int choice = Integer.parseInt(input.nextLine());

        switch (choice){
            case 1:
                account.displayBalance();
                break;

            case 2:
                System.out.println("Enter amount to withdraw");
                double amount = Double.parseDouble(input.nextLine());
                try {
                    account.withdraw(amount);
                }catch (RuntimeException e){
                    System.out.println(e.getMessage());
                }finally {
                    account.displayBalance();
                }
                break;

            case 3:
                System.out.println("Enter amount to deposit");
                amount = Double.parseDouble(input.nextLine());
                account.deposit(amount);
                account.displayBalance();
                break;

            case 4:
                System.out.println("Bye");
                System.exit(1);
                break;

            default:
                System.out.println("Invalid choice");
        }
    }
}
