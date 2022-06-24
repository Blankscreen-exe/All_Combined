from hammad.euler import euler as app
import streamlit as st

def eulerAPP():
    st.write("## Euler Formula Applet")
    x = st.slider("Select the value of X where X < Y:",
    min_value=-100, 
    max_value=100, 
    value=0)

    y = st.slider("Select the value of Y where X < Y:",
    min_value=-100, 
    max_value=100, 
    value=0)

    b = st.slider("Select the value of b(step-size):",
    min_value=0, 
    max_value=100, 
    value=3)
    
    # st.caption("Remeber not to let `X < Y` otherwise `None` will be returned")
    st.write("### Display Result:")
    st.write(f"### Distance =")
    st.write(app(x,y,b))