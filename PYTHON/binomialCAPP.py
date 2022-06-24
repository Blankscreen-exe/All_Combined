from haris.binomial_coefficient import binomial as app
import streamlit as st

def binomialCAPP():
    st.write("## Binomial Applet")
    x = st.slider("Select the value of X:",
    min_value=0, 
    max_value=100, 
    value=50)
    y = st.slider("Select the value of Y:",
    min_value=0, 
    max_value=100, 
    value=50)
    st.caption("Remeber not to let `X < Y` otherwise `None` will be returned")
    st.write("### Display Result:")
    st.write(f"### X = {app(x,y)}")