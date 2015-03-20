/*class MainFrame : contains the GUI methods*/
import javax.swing.*;
import javax.swing.UIManager.*;
import java.awt.*;
import java.awt.event.ActionListener;
import java.awt.event.ActionEvent;


public class MainFrame extends JFrame {

    private JButton enterButton;
    private JTextField inputTextField;
    private JLabel textLabel;

    public MainFrame(){
        super("Pig Latin");	//MainFrame window title
	

        try {
            UIManager.setLookAndFeel(UIManager.getSystemLookAndFeelClassName());
        } catch (Exception e) {
			/*use default JFrame look and feel*/
		}

        enterButton = new JButton("Enter");
        JPanel btnPanel = new JPanel(); //create a Panel for the enterButton
        btnPanel.add(enterButton);  //add enterButton to btnPanel

        inputTextField = new JTextField(20);
        JPanel textFieldPanel = new JPanel();	//create panel for text field
        textFieldPanel.add(inputTextField);	//add inputTextfield to textFieldPanel


        textLabel = new JLabel("Enter word to convert to Pig Latin");
        add(textLabel);

        setLayout(new BorderLayout());
        setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        add(btnPanel, BorderLayout.AFTER_LINE_ENDS);
        add(textFieldPanel, BorderLayout.CENTER);
        add(textLabel, BorderLayout.BEFORE_FIRST_LINE);
	
		//set window size and location based on the screen size
        Toolkit toolkit = Toolkit.getDefaultToolkit();
        Dimension d = toolkit.getScreenSize();
        setSize(d.width/4, d.height/6);
        setResizable(false);
        setLocation(d.width/3, d.height/3);

        setVisible(true);
	
        ButtonHandler handler = new ButtonHandler();
        enterButton.addActionListener( handler );
        inputTextField.addActionListener(handler);
    }

    private class ButtonHandler implements ActionListener {
        PigLatin hogLatin = new PigLatin();
        public void actionPerformed(ActionEvent event){
            String string = null;
	    //if enter button is pressed ?
            if(event.getSource() == enterButton || event.getSource() == inputTextField){
                string = hogLatin.toPigLatin(String.format(inputTextField.getText()));

                JOptionPane.showMessageDialog( null, string , "In Pig Latin", JOptionPane.PLAIN_MESSAGE);
            }
        }
    }
}

