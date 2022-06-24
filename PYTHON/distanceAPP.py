from neha.distance_zain import distance_formula as app
import streamlit as st

def distanceAPP():
    st.write("## Distance Formula Applet")
    x1 = st.slider("Select the value of X1:",
    min_value=-100, 
    max_value=100, 
    value=0)
    y1 = st.slider("Select the value of Y1:",
    min_value=-100, 
    max_value=100, 
    value=0)

    x2 = st.slider("Select the value of X2:",
    min_value=-100, 
    max_value=100, 
    value=0)
    y2 = st.slider("Select the value of Y2:",
    min_value=-100, 
    max_value=100, 
    value=0)
    
    # st.caption("Remeber not to let `X < Y` otherwise `None` will be returned")
    st.write("### Display Result:")
    st.write(f"### Distance = {app(x1,y1,x2,y2)}")