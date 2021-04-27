package main

// Categorical metrics
// let's say that we have a model that is supposed to predict
// some discrete value, such as fraud/not fraud, standing/sitting/walking,
// approved/not approved, and so on. Our data might look something like
// the following

/*
	observed, predicted
		0,		0
		0,		1
		2,		2
		1,		1
		1,		1
		0,		0
		2,		0
		0,		0
		..		.


	Undertanding these metrics and determine which is appropriate for
	our use case, we need to realiza that there a number of different scenarios
	that could occur when we are making discrete predictions




	True Positive(TP): We predicted a certain category, and the
						observation was actually that category
						(for example, we predicted fraud and the
						observation was fraud)


	False Positive(FP): We predicted a certain category, but te
						observation was actually another category
						(for example, we predicted fraud but the
						observation was not fraud)

	True Negative(TN): We predicted that the observation wasn't a
						certain category and the observation was
						not that category (for example, we predicted
						not fraud and the observation was not fraud)


	False Negative(FN): We predicted that the observation wasn't
						a certain category, but the observation was
						actually that category(for example, we predicted
						not fraud but the observation was fraud)



	You can see that there are a number of ways we can combine,
	aggregate, and measure these sceneraios. In fact, we could
	even aggregate/measeure theme in som sor of unique way related
	to our specifi ploblem. HJowever, there are some pretty standard
	ways of aggregating and measuring, these scenerarios that result
	in the fllowinf common metrics


*/

func main() {

}
