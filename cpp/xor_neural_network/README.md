# XOR Neural Network Simulation

XOR using neural network in C language.

## Description

This project shows that an array for XOR neural network simulation using the back-propagation learning algorithm. In input layer, there are 4 situations, which are 0^0=0, 0^1=1, 1^0=1, 1^1=0. We expected that the error of targets(desired outputs) and training output less than 0.01. Accuracy is defined as the training output should be less than 0.2 if the target is 0 and higher than 0.8 if the target is 1.

## Usage

```C
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <math.h>

#define LINES 4     //number of patterns or lines
#define INPUTS 3    //number of inputs for training set, includes two inputs and target
#define LINES_WTS 3 //number of lines for weights
#define INPUT_WTS 3 //number of weights per line
#define IN_FILE "XOR-data.txt"  //input data (i.e. patterns) file


#define NUMPAT 4 //number of patterns
#define NUMIN  3 //number of inputs (includes the bias)
#define NUMHID 4 //number of neurons in hidden layer
#define NUMOUT 1 //number of neurons in output layer
#define MIN -0.5
#define MAX 0.5
#define randn() (((double)rand()/((double)RAND_MAX + 1)) * (MAX - MIN)) + MIN

float input_data[LINES][INPUTS];
float weight_data[NUMHID][NUMIN];
float targets[LINES];


float Targets[NUMPAT][NUMOUT]; //same as targets array, but modified to work with network
float hidden[NUMPAT][NUMHID+1];
float deltaO[NUMOUT];
float weightsIH[NUMIN][NUMHID], weightsHO[NUMHID+1][NUMOUT];
float output[NUMPAT][NUMOUT];

//function for back propagation learning algorithm
void bp_learning(int p, float LearnRate, int numInputs, int numHidden, int numOutputs)
{
    int i, k;
    float sumDOW[numHidden+1], deltaH[numHidden+1];
    float deltaWeightIH[numInputs][numHidden], deltaWeightHO[numHidden+1][numOutputs];


    for (k = 0; k < numOutputs; k++) {
        for (i = 0; i < numHidden+1; i++) {
            deltaWeightHO[i][k] = 0.0;

        }
    }
    for (k = 0; k < numHidden; k++) {
        for (i = 0; i < numInputs; i++) {
            deltaWeightIH[i][k] = 0.0;

            }
        }

    for (i = 0; i < numHidden+1; i++) {
        sumDOW[i] = 0.0;
    }
    for (i = 0; i < numHidden+1; i++) {
        deltaH[i] = 0.0;
    }

    //back-propagate errors to hidden layer
    for (i = 0; i < numHidden+1; i++)  {
            for (k = 0; k < numOutputs; k++) {

                sumDOW[i] = sumDOW[i] + weightsHO[i][k] * deltaO[k];
        }

    }

        for (i = 1; i < numHidden+1; i++) {
            deltaH[i] = sumDOW[i] * hidden[p][i] * (1.0 - hidden[p][i]);

        }

    for (k = 0; k < numHidden; k++) {
        for (i = 0; i < numInputs; i++) {
            deltaWeightIH[i][k] = LearnRate * deltaH[k+1] * input_data[p][i]; //calculate weight change for weights going from input to hidden
            weightsIH[i][k] = weightsIH[i][k] + deltaWeightIH[i][k]; //update weights going from input to hidden

        }

    }
    for (k = 0; k < numOutputs; k++)  {
        for (i = 0; i < numHidden+1; i++) {
            deltaWeightHO[i][k] = LearnRate * deltaO[k] * hidden[p][i]; //calculate weight change for weights going from hidden to output
            weightsHO[i][k] = weightsHO[i][k] + deltaWeightHO[i][k]; //update weights going from hidden to output

        }
    }
}//end of bp_learning()

//function to report accuracy of network
float report_accuracy(int numPatterns, int numOutputs)
{
    float acc, out, hits = 0.0, misses = 0.0;
    int numPat = numPatterns, numOut = numOutputs;
    int p, k;

    for (k = 0; k < numOut; k++) {
        for (p = 0; p < numPat; p++) {

            if (output[p][k] <= 0.209999) {
                out = 0.0;
                if(Targets[p][k] == out) {
                    hits = hits + 1.0;
                    acc = (hits / (float)numPat) * 100.0;
                }
                else {
                    misses = misses + 1.0;
                    acc = (hits / (float)numPat) * 100.0;
                }
            }
            else if (output[p][k] >= 0.8) {
                out = 1.0;
                if (Targets[p][k] == out) {
                    hits = hits + 1.0;
                    acc = (hits / (float)numPat) * 100.0;
                }
                else {
                    misses = misses + 1.0;
                    acc = (hits / (float)numPat) * 100.0;
                }
            }
             else {
                misses = misses + 1.0; //keep track of misses
                acc = (hits / (float)numPat) * 100.0; //accuracy of network in percentage of hits
            }
        }
    }
    return acc;
}

//function to simulate network
void simulate_net()
{
    int i, k, p, epoch;
    int numPatterns = NUMPAT, numInputs = NUMIN, numHidden = NUMHID, numOutputs = NUMOUT;
    float error, LearnRate = 0.7, acc;
    float sumH[numPatterns][numHidden];
    float sumO[numPatterns][numOutputs];


    for (k = 0; k < numHidden; k++) {
        for (i = 0; i < numInputs; i++) {
            //deltaWeightIH[i][k] = 0.0;
            weightsIH[i][k] = randn();  //random weights for input to hidden layer
            printf("\nweightsIH[%d][%d]: %f", i, k, weightsIH[i][k]);
            }
        }

    for (k = 0; k < numOutputs; k++) {
        for (i = 0; i < numInputs; i++) {
            weightsHO[i][k] = randn();  //random weights for hidden layer to output layer
            printf("\nweightsHO[%d][%d]: %f", i, k, weightsHO[i][k]);
        }
    }

    for (k = 0; k < numOutputs; k++) {
        for (i = 0; i < numPatterns; i++) {
            Targets[i][k] = targets[i];
            printf("\n\nTargets[%d][%d]: %f", i, k, Targets[i][k]);
        }
    }

    for (epoch = 1; epoch <= 3000; epoch++) {
        error = 0.0;
        for (p = 0; p < numPatterns; p++) { //repeat for all training patterns

            for (k = 0; k < numHidden; k++) {
                sumH[p][k] = 0.0;
                for (i = 0; i < numInputs; i++) {

                    sumH[p][k] = sumH[p][k] + input_data[p][i] * weightsIH[i][k]; //sum of outputs at hidden layer, bias included in input_data
                }

                if (k == 0) {
                    hidden[p][k] = 1.0; //add bias at hidden layer
                }
                hidden[p][k+1] = 1.0 / (1.0 + exp(-sumH[p][k]));    //sigmoidal outputs at hidden layer

            }
            for (k = 0; k < numOutputs; k++) {
                sumO[p][k] = 0.0;
                for (i = 0; i < numHidden+1; i++) {
                    sumO[p][k] = sumO[p][k] + hidden[p][i] * weightsHO[i][k]; //sum of outputs at output layer
                }

                output[p][k] = 1.0 / (1.0 + exp(-sumO[p][k]));    //sigmoidal output at output layer

                error = error + 0.5 * (Targets[p][k] - output[p][k]) * (Targets[p][k] - output[p][k]); //calculate error

                deltaO[k] = (Targets[p][k] - output[p][k]) * output[p][k] * (1.0 - output[p][k]);

            }
            printf("\nepoch %-5d :  Error = %f", epoch, error);
            bp_learning(p, LearnRate, numInputs, numHidden, numOutputs); //function for back propagation learning algorithm
        }//end of training patterns

        if (error < 0.05) { //stop learning if error is less than 0.05
            break;
        }

    }//end of epoch

    acc = report_accuracy(numPatterns, numOutputs);

    printf("\n\n\tNetwork Results\n\nPat\t");

    for (i = 1; i < numInputs; i++) {
        printf("Input%-4d\t", i);
    }
    for (k = 1; k <= numOutputs; k++) {
        printf("Targets\t\tOutputs\t");
    }
    for (p = 0; p < numPatterns; p++) {
        printf("\n%d\t", p);
        for (i = 1; i < numInputs; i++) {
            printf("%f\t", input_data[p][i]);
        }
        for(k = 0; k < numOutputs; k++) {
            printf("%f\t%f\t", Targets[p][k], output[p][k]);
        }
    }
    printf("\n\n\tAccuracy = %f%%\n", acc);


}//end of simulate_net()


// function to read XOR data file contents
void read_input_data()
{
    FILE *fp;
    fp = fopen(IN_FILE, "r");   //open text file

    if (fp == NULL) {
        printf("Error while opening the inputs data file.\n");
        return;
    }
    else {
    int i,k;
    char read_in[150];
    float conv_in;

    fscanf(fp, "%*[^\n]");      //skip first line of text file
    if (LINES == 4) {
        for (i = 0; i < LINES; i++){
            for (k = 0; k <= INPUTS; k++){

            //add bias to all patterns (i.e. lines)

            if (k == 0) {
                input_data[i][k] = 1;           //add bias to first element of each pattern
            }

            // store targets
            else if (k == 3) {
                fscanf(fp, "%s", read_in);      //read single string up to the whitespace
                conv_in = atof(read_in);        //convert from character to floating point integer
                targets[i] = conv_in;
            }
            else {
                fscanf(fp, "%s", read_in);      //read single string up to the whitespace
                conv_in = atof(read_in);        //convert from character to floating point integer
                input_data[i][k] = conv_in;
                }
            }

        }
        fclose(fp);
    }
    else {
        printf("Number of patterns (i.e. lines) do not match that of the text file\n");
        }

    }
}

int main()
{

    read_input_data();        //call function to read XOR data file contents
    simulate_net();           //call function to simulate network
    return 0;

}
```

## Cautions

1. Shall have XORTarget.txt.

## Operation

- ./main
