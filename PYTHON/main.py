import streamlit as st

from binomialCAPP import binomialCAPP as BC
from distanceAPP import distanceAPP as DIST
from eulerAPP import eulerAPP as EULER

options = [
    "binomial",
    "distance",
    "euler",
    "euler identity",
    "fibonacci",
    "fourier",
    "gaussian",
    "laplace",
    "prime numbers",
    "quadratic formula",
    "relativity",
    "slope",
    "standard deviation",
    "sum of numbers",
    "taylor series",
    "trigonometry"
]

st.title("Math Formula App")

app_select = st.selectbox("Select an Applet:", options, key="app-select")

if app_select == options[0]: #binomial
    BC()
elif app_select == options[1]: #distance
    DIST()
elif app_select == options[2]: #euler
    EULER()
elif app_select == options[3]:
    pass
elif app_select == options[4]:
    pass
elif app_select == options[5]:
    pass
elif app_select == options[6]:
    pass
elif app_select == options[7]:
    pass
elif app_select == options[8]:
    pass
elif app_select == options[9]:
    pass
elif app_select == options[10]:
    pass
elif app_select == options[11]:
    pass
elif app_select == options[12]:
    pass
elif app_select == options[13]:
    pass
elif app_select == options[14]:
    pass
elif app_select == options[15]:
    pass
elif app_select == options[16]:
    pass
elif app_select == options[17]:
    pass