package main

/*
Confusion matrices, AUC, and ROC

In addition to calculatin individual numerical metrics for our
models, there are a varietu of techiques to combine various metrics
into a form that gives you a more complete representation of model
perfomance. These include, but are certainly not limite to,
consusion matrices and
area under the curve(AUC)/Receiver Operating Characteristic (ROC)curves

Confision matrices allow us to visualize the varous TP, TN, FP and FN
values that we predict in a two-dimensional format. A confusion matrix
has rows corresponding to the categories that you ere supposed to predict,
and columns corresponding to categories that were predicted.
Then, the value of each element is the corresponding count:

								Predicted
						Fraud		Not fraud
Observed	fraud		  TP			FN
			Not fraund	  FP			TN

*/

func main() {

}
