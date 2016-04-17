package com.pitswarner;

import android.Manifest;
import android.content.Context;
import android.content.pm.PackageManager;
import android.hardware.Sensor;
import android.hardware.SensorEvent;
import android.hardware.SensorEventListener;
import android.hardware.SensorManager;
import android.location.Location;
import android.location.LocationListener;
import android.location.LocationManager;
import android.os.Bundle;
import android.os.Environment;
import android.support.v4.app.ActivityCompat;
import android.support.v7.app.AppCompatActivity;
import android.util.Log;
import android.widget.TextView;
import android.widget.Toast;

import java.io.DataInputStream;
import java.io.DataOutputStream;
import java.io.File;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.FileWriter;
import java.io.IOException;
import java.io.InputStream;
import java.net.HttpURLConnection;
import java.net.MalformedURLException;
import java.net.URL;
import java.util.Timer;
import java.util.TimerTask;

import butterknife.Bind;
import butterknife.ButterKnife;

public class MainActivity extends AppCompatActivity implements SensorEventListener {
    @Bind(R.id.timestamp)
    TextView timestampView;
    @Bind(R.id.x)
    TextView xView;
    @Bind(R.id.y)
    TextView yView;
    @Bind(R.id.z)
    TextView zView;
    @Bind(R.id.acc)
    TextView accView;
    @Bind(R.id.bear)
    TextView bearView;
    @Bind(R.id.speed)
    TextView speedView;
    @Bind(R.id.lng)
    TextView lngView;
    @Bind(R.id.lat)
    TextView latView;
    @Bind(R.id.alt)
    TextView altView;


    private SensorManager sensorManager;
    private Sensor accelerometer;

    private LocationManager locationManager;
    private LocationListener locationListener = new LocationListener() {
        @Override
        public void onLocationChanged(Location location) {
            synchronized (locationSyncRoot) {
                MainActivity.this.location = new Location(location);
            }
        }

        public void onStatusChanged(String provider, int status, Bundle extras) {
            Toast.makeText(MainActivity.this, provider + " " + status, Toast.LENGTH_SHORT).show();
        }

        public void onProviderEnabled(String provider) {
            if (ActivityCompat.checkSelfPermission(MainActivity.this, Manifest.permission.ACCESS_FINE_LOCATION) != PackageManager.PERMISSION_GRANTED && ActivityCompat.checkSelfPermission(MainActivity.this, Manifest.permission.ACCESS_COARSE_LOCATION) != PackageManager.PERMISSION_GRANTED) {
                return;
            }
            location = locationManager.getLastKnownLocation(LocationManager.GPS_PROVIDER);
            Toast.makeText(MainActivity.this, provider + " enabled", Toast.LENGTH_SHORT).show();
        }

        public void onProviderDisabled(String provider) {
            Toast.makeText(MainActivity.this, provider + " disabled", Toast.LENGTH_SHORT).show();
        }
    };


    private Timer timer = new Timer();
    private TimerTask collectTask = new TimerTask() {
        @Override
        public void run() {
            final SensorEvent event;
            final Location location;
            synchronized (sensorSyncRoot) {
                event = lastEvent;
            }

            synchronized (locationSyncRoot) {
                location = MainActivity.this.location;
            }

            final long timestamp = event == null ? 0 : event.timestamp / 1000L;
            final float x = event == null ? 0 : event.values[0];
            final float y = event == null ? 0 : event.values[1];
            final float z = event == null ? 0 : event.values[2];
            final float acc = location == null ? 0 : location.getAccuracy();
            final float bear = location == null ? 0 : location.getBearing();
            final float speed = location == null ? 0 : location.getSpeed();
            final double lng = location == null ? 0 : location.getLongitude();
            final double lat = location == null ? 0 : location.getLatitude();
            final double alt = location == null ? 0 : location.getAltitude();

            synchronized (fileSyncRoot) {
                try {
                    writer.write(
                            "\n" + timestamp + ',' +
                            x + ',' + y + ',' + z + ',' +
                            acc + ',' + bear + ',' + speed + ',' +
                            lng + ',' + lat + ',' + alt);
                } catch (IOException e) {
                    Log.e("pizdec", "write", e);
                }
            }


            runOnUiThread(new Runnable() {
                @Override
                public void run() {
                    timestampView.setText(String.valueOf(timestamp));
                    xView.setText(String.valueOf(x));
                    yView.setText(String.valueOf(y));
                    zView.setText(String.valueOf(z));
                    accView.setText(String.valueOf(acc));
                    bearView.setText(String.valueOf(bear));
                    speedView.setText(String.valueOf(speed));
                    lngView.setText(String.valueOf(lng));
                    latView.setText(String.valueOf(lat));
                    altView.setText(String.valueOf(alt));
                }
            });
        }
    };
    private TimerTask sendDataTask = new TimerTask() {
        @Override
        public void run() {
            File file;

            synchronized (fileSyncRoot) {
                file = dataFile;
                try {
                    writer.flush();
                    writer.close();
                    createNewFile();
                } catch (IOException e) {
                    Log.e("pizdec", "closewriter", e);
                }
            }

            HttpURLConnection conn = null;
            DataOutputStream dos = null;
            String lineEnd = "\r\n";
            String twoHyphens = "--";
            String boundary = "*****";
            int bytesRead, bytesAvailable, bufferSize;
            byte[] buffer;
            int maxBufferSize = 1 * 1024 * 1024;
            File sourceFile = file;

            try {

                // open a URL connection to the Servlet
                FileInputStream fileInputStream = new FileInputStream(sourceFile);
                URL url = new URL("http://52.58.116.75:8080/measures");

                // Open a HTTP  connection to  the URL
                conn = (HttpURLConnection) url.openConnection();
                conn.setDoInput(true); // Allow Inputs
                conn.setDoOutput(true); // Allow Outputs
                conn.setUseCaches(false); // Don't use a Cached Copy
                conn.setRequestMethod("POST");
                conn.setRequestProperty("Connection", "Keep-Alive");
                conn.setRequestProperty("ENCTYPE", "multipart/form-data");
                conn.setRequestProperty("Content-Type", "multipart/form-data;boundary=" + boundary);
                conn.setRequestProperty("uploaded_file", file.getName());

                dos = new DataOutputStream(conn.getOutputStream());

                dos.writeBytes(twoHyphens + boundary + lineEnd);
                dos.writeBytes("Content-Disposition: form-data; name=\"uploadfile\";filename=\"" + file.getName() + "\"" + lineEnd);
                dos.writeBytes(lineEnd);

                // create a buffer of  maximum size
                bytesAvailable = fileInputStream.available();

                bufferSize = Math.min(bytesAvailable, maxBufferSize);
                buffer = new byte[bufferSize];

                // read file and write it into form...
                bytesRead = fileInputStream.read(buffer, 0, bufferSize);

                while (bytesRead > 0) {

                    dos.write(buffer, 0, bufferSize);
                    bytesAvailable = fileInputStream.available();
                    bufferSize = Math.min(bytesAvailable, maxBufferSize);
                    bytesRead = fileInputStream.read(buffer, 0, bufferSize);

                }

                // send multipart form data necesssary after file data...
                dos.writeBytes(lineEnd);
                dos.writeBytes(twoHyphens + boundary + twoHyphens + lineEnd);

                InputStream is = new DataInputStream(conn.getInputStream());

                Log.v("tag", String.valueOf(is.available()));


//                // Responses from the server (code and message)
//                int serverResponseCode = conn.getResponseCode();
//                String serverResponseMessage = conn.getResponseMessage();
//
//                Log.i("uploadFile", "HTTP Response is : "
//                        + serverResponseMessage + ": " + serverResponseCode);

                //close the streams //
                fileInputStream.close();
                dos.close();

            } catch (MalformedURLException ex) {
                Log.e("Upload file to server", "error: " + ex.getMessage(), ex);
            } catch (Exception e) {
                Log.e("Upload file to server", "Exception", e);
            }


//            HttpURLConnection conn;
//            DataOutputStream dos;
//            String lineEnd = "\r\n";
//            String twoHyphens = "--";
//            String boundary = "*****";
//            int bytesRead;
//            byte[] buffer= new byte[1000];
//
//            try {
//                // open a URL connection to the Servlet
//                FileInputStream fileInputStream = new FileInputStream(file);
//                URL url = new URL("http://52.58.116.75:8080/measures");
//
//                // Open a HTTP  connection to  the URL
//                conn = (HttpURLConnection) url.openConnection();
//                conn.setDoInput(true); // Allow Inputs
//                conn.setDoOutput(true); // Allow Outputs
//                conn.setUseCaches(false); // Don't use a Cached Copy
//                conn.setRequestMethod("POST");
//                conn.setRequestProperty("Connection", "Keep-Alive");
//                conn.setRequestProperty("ENCTYPE", "multipart/form-data");
//                conn.setRequestProperty("Content-Type", "multipart/form-data;boundary=" + boundary);
//
//                dos = new DataOutputStream(conn.getOutputStream());
//
//                dos.writeBytes(twoHyphens + boundary + lineEnd);
//                dos.writeBytes("Content-Disposition: form-data; name=\"uploadfile\";filename=\"" + file.getName() + "\"" + lineEnd);
//                dos.writeBytes("Content-Type: application/octet-stream");
//                dos.writeBytes(lineEnd);
//
//                do {
//                    bytesRead = fileInputStream.read(buffer, 0, buffer.length);
//                    if (bytesRead > 0)
//                        dos.write(buffer, 0, bytesRead);
//                }  while (bytesRead > 0);
//
//                // send multipart form data necesssary after file data...
//                dos.writeBytes(lineEnd);
//                dos.writeBytes(twoHyphens + boundary + twoHyphens + lineEnd);
//                dos.close();
//            } catch (IOException e) {
//                Log.e("pizdec", "send", e);
//            }
        }
    };

    private final  Object sensorSyncRoot = new Object();
    private SensorEvent lastEvent = null;

    private final Object locationSyncRoot = new Object();
    private Location location = null;

    private final Object fileSyncRoot = new Object();
    private File dataFile = null;
    private FileWriter writer = null;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        ButterKnife.bind(this);

        sensorManager = (SensorManager) getSystemService(Context.SENSOR_SERVICE);
        accelerometer = sensorManager.getDefaultSensor(Sensor.TYPE_ACCELEROMETER);

        locationManager = (LocationManager) this.getSystemService(Context.LOCATION_SERVICE);
    }

    @Override
    protected void onResume() {
        super.onResume();

        try {
            createNewFile();
        } catch (FileNotFoundException e) {
            Log.e("pizdec", "stream", e);
        } catch (IOException e) {
            Log.e("pizdec", "create", e);
        }

        if (ActivityCompat.checkSelfPermission(this, Manifest.permission.ACCESS_FINE_LOCATION) != PackageManager.PERMISSION_GRANTED &&
                ActivityCompat.checkSelfPermission(this, Manifest.permission.ACCESS_COARSE_LOCATION) != PackageManager.PERMISSION_GRANTED)
            return;
        location = locationManager.getLastKnownLocation(LocationManager.GPS_PROVIDER);
        locationManager.requestLocationUpdates(LocationManager.GPS_PROVIDER, 0, 0, locationListener);
        sensorManager.registerListener(this, accelerometer, 100);
        timer.schedule(collectTask, 1000, 100);
        timer.schedule(sendDataTask, 10000, 60000);
    }

    private void createNewFile() throws IOException {
        synchronized (fileSyncRoot) {
            dataFile = new File(Environment.getExternalStoragePublicDirectory(Environment.DIRECTORY_DOWNLOADS), "pitsdata");
            if (dataFile.mkdirs())
                Log.v("verboze", dataFile.getAbsolutePath());
            dataFile = new File(dataFile, System.currentTimeMillis() + ".csv");
            if (dataFile.createNewFile())
                Log.v("verboze", dataFile.getAbsolutePath());
            writer = new FileWriter(dataFile);
            writer.write("timestamp,acceleration.x,acceleration.y,acceleration.z,accuracy,bearing,speed,longitude,latitude,altitude");
        }
    }

    @Override
    protected void onPause() {
        super.onPause();
        collectTask.cancel();
        sendDataTask.cancel();
        sensorManager.unregisterListener(this, accelerometer);
        if (ActivityCompat.checkSelfPermission(this, Manifest.permission.ACCESS_FINE_LOCATION) != PackageManager.PERMISSION_GRANTED &&
                ActivityCompat.checkSelfPermission(this, Manifest.permission.ACCESS_COARSE_LOCATION) != PackageManager.PERMISSION_GRANTED)
            return;

        locationManager.removeUpdates(locationListener);
        try {
            writer.close();
        } catch (IOException e) {
            Log.e("pizdec", "close", e);
        }
    }

    @Override
    public void onSensorChanged(SensorEvent event) {
        synchronized (sensorSyncRoot) {
            lastEvent = event;
        }
    }

    @Override
    public void onAccuracyChanged(Sensor sensor, int accuracy) { }
}
