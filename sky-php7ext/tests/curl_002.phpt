--TEST--
skywalking extension is available
--CREDITS--

--SKIPIF--
<?php if (!extension_loaded("skywalking")) print "skip"; ?>
--FILE--
<?php
  
  $logPath = ini_get("skywalking.log_path");
  $logFilename = strtolower($logPath . DIRECTORY_SEPARATOR .  'skywalking.' . date("Ymd") . '.log');
  echo "Please use the command( tail -f *.log ) to view\n";
  include 'server.inc';
  $host = curl_cli_server_start();

  // start testing
  echo '*** Testing curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1); ***' . "\n";

  $url = "{$host}/get.php?test=get";
  $ch = curl_init();

  ob_start(); // start output buffering
  curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
  curl_setopt($ch, CURLOPT_URL, $url); //set the url we want to use
  
  $curl_content = curl_exec($ch);
  curl_close($ch);

  var_dump( $curl_content );
?>
===DONE===
--EXPECTF--
Please use the command( tail -f *.log ) to view
*** Testing curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1); ***
string(25) "Hello World!
Hello World!"
===DONE===
