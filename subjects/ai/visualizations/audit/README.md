#### Exercise 0: Environment and libraries

##### The exercice is validated is all questions of the exercice are validated

##### Activate the virtual environment. If you used `conda` run `conda activate your_env`

##### Run `python --version`

###### Does it print `Python 3.x`? x >= 8

##### Does `import jupyter`, `import numpy`, `import pandas`, `matplotlib` and `plotly` run without any error ?

---

---

#### Exercise 1: Pandas plot 1

##### The exercice is validated is all questions of the exercice are validated

##### The solution of question 1 is accepted if the plot reproduces the plot in the image and respect those criteria.

###### Does it have a the title ?

###### Does it have a name on x-axis and y-axis ?

![alt text][logo_ex6]

[logo_ex6]: ../w1day03_ex6_plot1.png "Time series ex6"

##### The solution of question 2 is accepted if the plot reproduces the plot in the image by using `plotly.graph_objects` and respect those criteria.

2.This question is validated if the plot is in the image is reproduced using `plotly.graph_objects` given those criteria:

###### Does it have a the title ?

###### Does it have a name on x-axis and y-axis ?

![alt text][logo_ex6]

[logo_ex6]: ../w1day03_ex6_plot1.png "Time series ex6"

---

---

#### Exercise 2: Pandas plot 2

##### The solution of question 1 is accepted if the plot reproduces the plot in the image and respect the following criteria. It is important to observe that the older people are, the the more children they have.

###### Does it have a the title ?

###### Does it have a name on x-axis and y-axis ?

![alt text][logo_ex2]

[logo_ex2]: ../w1day03_ex2_plot1.png "Scatter plot ex2"

---

---

#### Exercise 3: Matplotlib 1

##### The solution of question 1 is accepted if the plot reproduces the plot in the image and respect those criteria.

###### Does it have a the title ?

###### Does it have a name on x-axis and y-axis ?

###### Are the x-axis and y-axis limited to [1,8] ?

###### Is the line a red dashdot line with a width of 3 ?

###### Are the circles blue circles with a size of 12 ?

![alt text][logo_ex3]

[logo_ex3]: ../w1day03_ex3_plot1.png "Scatter plot ex3"

---

---

#### Exercise 4: Matplotlib 2

##### The solution of question 1 is accepted if the plot reproduces the plot in the image and respect those criteria.

###### Does it have a the title ?

###### Does it have a name on x-axis and y-axis ?

###### Is the left data black ?

###### Is the right data red ?

![alt text][logo_ex4]

[logo_ex4]: ../w1day03_ex4_plot1.png "Twin axis ex4"

https://matplotlib.org/gallery/api/two_scales.html

---

---

#### Exercise 5: Matplotlib subplots

##### The solution of question 1 is accepted if the plot reproduces the plot in the image and respect those criteria.

###### Does it contain 6 subplots (2 rows, 3 columns)?

###### Does it have space between plots (`hspace=0.5` and `wspace=0.5`)?

###### Do all subplots contain a title: `Title i` ?

###### Do all subplots contain a text `(2,3,i)` centered at `(0.5, 0.5)`? _Hint_: check the parameter `ha` of `text`

###### Have all subplots been created in a for loop ?

![alt text][logo_ex5]

[logo_ex5]: ../w1day03_ex5_plot1.png "Subplots ex5"

---

---

#### Exercise 6: Plotly 1

##### The exercice is validated is all questions of the exercice are validated

##### The solution of question 1 is accepted if the plot reproduces the plot in the image and respect those criteria.

###### Does it have a the title ?

###### Does it have a name on x-axis and y-axis ?

![alt text][logo_ex6]

[logo_ex6]: ../w1day03_ex6_plot1.png "Time series ex6"

##### The solution of question 2 is accepted if the plot reproduces the plot in the image by using `plotly.graph_objects` and respect those criteria.

2.This question is validated if the plot is in the image is reproduced using `plotly.graph_objects` given those criteria:

###### Does it have a the title ?

###### Does it have a name on x-axis and y-axis ?

![alt text][logo_ex6]

[logo_ex6]: ../w1day03_ex6_plot1.png "Time series ex6"

---

---

#### Exercise 7: Plotly Box plots

##### The solution of question 1 is accepted if the plot reproduces the plot in the image and respect those criteria. The code below shows a solution.

###### Does it have a the title ?

###### Does it have a legend ?

![alt text][logo_ex7]

[logo_ex7]: ../w1day03_ex7_plot1.png "Box plot ex7"

```python
import plotly.graph_objects as go
import numpy as np

y0 = np.random.randn(50)
y1 = np.random.randn(50) + 1 # shift mean
y2 = np.random.randn(50) + 2

fig = go.Figure()
fig.add_trace(go.Box(y=y0, name='Sample A',
                marker_color = 'indianred'))
fig.add_trace(go.Box(y=y1, name = 'Sample B',
                marker_color = 'lightseagreen'))

fig.show()
```
