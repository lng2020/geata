package handler

// IEC61850Client take ip and port arguments, then it uses those argument to invoke an execute file.
// The execute file periodly return string streams and if the IEC61850 client recevive the string,
// it will parse the string and send back to IEC61850handler
