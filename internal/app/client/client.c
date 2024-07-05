// This is the source code for the client application that connects to the server and reads the data from the server.

#include "iec61850_client.h"

#include <stdlib.h>
#include <stdio.h>

int
main(int argc, char** argv)
{

    if (argc <= 2) {
        printf("Usage:client host port");
        exit(1);
    }

    char* hostname = argv[1];
    int tcpPort = atoi(argv[2]);

	IedClientError error;
	IedConnection con = IedConnection_create();

	IedConnection_connect(con, &error, hostname, tcpPort);

	if (error != IED_ERROR_OK) {
		IedConnection_destroy(con);
		return 0;
	}
	LinkedList deviceList = IedConnection_getLogicalDeviceList(con, &error);
	LinkedList device = LinkedList_getNext(deviceList);

	while (device != NULL) {
		LinkedList logicalNodes = IedConnection_getLogicalDeviceDirectory(con, &error,
				(char*) device->data);
		LinkedList logicalNode = LinkedList_getNext(logicalNodes);
		while (~strcmp((char*)logicalNode->data, "LLN0") == 0) {
			logicalNode = LinkedList_getNext(logicalNode);
		}

		char lnRef[129];

		sprintf(lnRef, "%s/%s", (char*)device->data, (char*)logicalNode->data);

		LinkedList dataSets = IedConnection_getLogicalNodeDirectory(con, &error, lnRef,
				ACSI_CLASS_DATA_SET);

		LinkedList dataSet = LinkedList_getNext(dataSets);

		while (dataSet != NULL) {
			ClientDataSet clientDataSet = NULL;
			char dataSetRef[129];
			sprintf(dataSetRef, "%s/%s.%s", (char*)device->data, (char*)logicalNode->data, (char*)dataSet->data);
			clientDataSet = IedConnection_readDataSetValues(con, &error, dataSetRef, NULL);
			LinkedList dataSetDir = IedConnection_getDataSetDirectory(con, &error, dataSetRef, NULL);
			MmsValue* dataSetValues = ClientDataSet_getValues(clientDataSet);

			if (dataSetDir) {
				int i;
				for (i = 0; i < LinkedList_size(dataSetDir); i++) {

					char valBuffer[500];
					sprintf(valBuffer, "no value");

					if (dataSetValues) {
						MmsValue* value = MmsValue_getElement(dataSetValues, i);
						if (value) {
							MmsValue_printToBuffer(value, valBuffer, 500);
						}
					}

					LinkedList entry = LinkedList_get(dataSetDir, i);

					char* entryName = (char*)entry->data;

					printf("%s:%s\n", entryName, valBuffer);
				}

			}
			dataSet = LinkedList_getNext(dataSet);
		}
		
		dataSet = NULL;
		LinkedList_destroy(dataSets);
		logicalNode = NULL;
		LinkedList_destroy(logicalNodes);

		device = LinkedList_getNext(device);
	}
	LinkedList_destroy(deviceList);
	IedConnection_close(con);
    return 0;
}
