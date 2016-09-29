package md01.cordova.joeblew.org.md01;

import android.os.Bundle;
import android.support.design.widget.FloatingActionButton;
import android.support.design.widget.Snackbar;
import android.support.v7.app.AppCompatActivity;
import android.support.v7.widget.Toolbar;
import android.view.View;
import android.view.Menu;
import android.view.MenuItem;


// import go.bind.Bind;

// import go.hello.Hello;

import go.logpackermobilesdk.Logpackermobilesdk;
//import go.cli.Cli;



public class MainActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        Toolbar toolbar = (Toolbar) findViewById(R.id.toolbar);
        setSupportActionBar(toolbar);


        FloatingActionButton fab = (FloatingActionButton) findViewById(R.id.fab);

        fab.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {


                String textHello = "empty :)";
                try {
                    textHello = "new shit ";

                } catch (Exception e) {
                    e.printStackTrace();
                }

                String logPackerText = "nul";
                try {
                    /*
                    Logpackermobilesdk.Client client = Logpackermobilesdk.newClient("https://logpacker.mywebsite.com", "dev", android.os.Build.MODEL);

                    Logpackermobilesdk.Message msg = client.NewMessage();
                    msg.setMessage("Crash is here!");
                    // Use another optional setters for msg object

                    client.Send(msg); // Send will return Cluster response
                    */

                    logPackerText = "sent";
                } catch (Exception e) {
                    // Cannot connect to Cluster or validation error
                    logPackerText = e.toString();
                }

                /*
                String textHello = Hello.Greetings("Android and Gopher");
                */

                Snackbar.make(view, "Replace with your own action" + logPackerText, Snackbar.LENGTH_LONG)
                        .setAction("Action", null).show();
            }
        });

    }

    @Override
    public boolean onCreateOptionsMenu(Menu menu) {
        // Inflate the menu; this adds items to the action bar if it is present.
        getMenuInflater().inflate(R.menu.menu_main, menu);
        return true;
    }

    @Override
    public boolean onOptionsItemSelected(MenuItem item) {
        // Handle action bar item clicks here. The action bar will
        // automatically handle clicks on the Home/Up button, so long
        // as you specify a parent activity in AndroidManifest.xml.
        int id = item.getItemId();

        //noinspection SimplifiableIfStatement
        if (id == R.id.action_settings) {
            return true;
        }

        return super.onOptionsItemSelected(item);
    }
}
