import pandas as pd
import matplotlib.pyplot as plt
import numpy as np

# Load the data from an Excel file
try:
    # Use read_excel for .xlsx files
    data = pd.read_excel('abc.xlsx', engine='openpyxl')  # Adjust the filename as needed

    # Print raw data for inspection
    print("Raw Data:")
    print(data)

    # Print column names for debugging
    print("Column Names:", data.columns)

    # Check for missing values
    if data.isnull().values.any():
        print("Warning: There are missing values in the dataset.")

    # Check if longitude values are within valid range
    if not all((-180 <= x <= 180 for x in data['Longitude_Start'].dropna())):
        print("Warning: Some longitude values are out of valid range (-180 to 180).")

    # Check if latitude values are within valid range
    if not all((-90 <= y <= 90 for y in data['Latitude_Start'].dropna())):
        print("Warning: Some latitude values are out of valid range (-90 to 90).")

    # Initialize lists for plotting
    start_longitudes = []
    start_latitudes = []
    stop_longitudes = []
    stop_latitudes = []

    # Loop through the DataFrame to prepare data for plotting
    for index, row in data.iterrows():
        start_longitudes.append(float(row['Longitude_Start']))
        start_latitudes.append(float(row['Latitude_Start']))
        stop_longitudes.append(float(row['Longitude_Stop']))
        stop_latitudes.append(float(row['Latitude_Stop']))

    # Create scatter plot for breeding areas (Longitude_Start vs Latitude_Start)
    plt.figure(figsize=(10, 6))

    # Plot breeding area points with different colors based on Population
    for i in range(len(start_longitudes)):
        if data['Population'][i] == 'Chaun':
            plt.scatter(start_longitudes[i], start_latitudes[i], marker='o', color='purple')
        elif data['Population'][i] == 'Indigirka':
            plt.scatter(start_longitudes[i], start_latitudes[i], marker='^', color='blue')

    # Remove the plotting of wintering area points (red 'x')
    # for i in range(len(stop_longitudes)):
    #     plt.scatter(stop_longitudes[i], stop_latitudes[i], marker='x', color='red')

    # Define the slope and intercept from the given equation
    slope = 1.2919
    intercept = -10.107

    # Generate x values from 100 to 180 with a step of 10 (for breeding areas)
    x_values = np.arange(100, 181, 10)  # [100, 110, ..., 180]

    # Calculate corresponding y values using the regression equation for x_values
    regression_y_values = slope * x_values + intercept

    # Plot the regression line using x_values and corresponding regression_y_values
    plt.plot(x_values, regression_y_values, color='green', label='Regression Line')

    # Set y-axis ticks from 0 to 90 with a step of 10
    plt.yticks(np.arange(0, 91, 10))
    
    # Set x-axis ticks from 100 to 180 with a step of 10
    plt.xticks(np.arange(100, 181, 10))

    # Expand the axes limits
    plt.xlim(100, 180)   # X-axis limits from 100 to 180
    plt.ylim(0, 90)      # Y-axis limits from 0 to 90

    # Add labels and title
    plt.xlabel('Longitude (Breeding Area)')
    plt.ylabel('Longitude (Wintering Area)')

    # Add R² value and slope-intercept equation text annotations in the top left corner of the graph
    r_squared = 0.6624
    plt.text(110, 85, f'R² = {r_squared:.4f}', fontsize=10, verticalalignment='top', horizontalalignment='left')
    plt.text(110, 80, f'y = {slope:.4f}x + {intercept:.3f}', fontsize=10, verticalalignment='top', horizontalalignment='left')

    # Remove grid lines by commenting out or removing this line:
    # plt.grid()

    # Show the plot without a legend
    plt.show()

except Exception as e:
    print("Error reading the Excel file:", e)
