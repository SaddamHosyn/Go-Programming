# Load necessary libraries
library(readxl)
library(ggplot2)

# Load the data from an Excel file
tryCatch({
  # Read the Excel file (adjust the filename as needed)
  data <- read_excel('abc.xlsx')

  # Print raw data for inspection
  cat("Raw Data:\n")
  print(data)

  # Print column names for debugging
  cat("Column Names:\n", colnames(data), "\n")

  # Check for missing values
  if (any(is.na(data))) {
    cat("Warning: There are missing values in the dataset.\n")
  }

  # Check if longitude values are within the valid range
  if (any(data$Longitude_Start < -180 | data$Longitude_Start > 180, na.rm = TRUE)) {
    cat("Warning: Some longitude values are out of valid range (-180 to 180).\n")
  }

  # Check if latitude values are within the valid range
  if (any(data$Latitude_Start < -90 | data$Latitude_Start > 90, na.rm = TRUE)) {
    cat("Warning: Some latitude values are out of valid range (-90 to 90).\n")
  }

  # Initialize lists for plotting
  start_longitudes <- numeric()
  start_latitudes <- numeric()
  stop_longitudes <- numeric()
  stop_latitudes <- numeric()

  # Loop through the data to prepare for plotting
  for (i in 1:nrow(data)) {
    start_longitudes <- c(start_longitudes, as.numeric(data$Longitude_Start[i]))
    start_latitudes <- c(start_latitudes, as.numeric(data$Latitude_Start[i]))
    stop_longitudes <- c(stop_longitudes, as.numeric(data$Longitude_Stop[i]))
    stop_latitudes <- c(stop_latitudes, as.numeric(data$Latitude_Stop[i]))
  }

  # Create a data frame for plotting breeding areas
  plot_data <- data.frame(
    Longitude_Start = start_longitudes,
    Latitude_Start = start_latitudes,
    Population = data$Population
  )

  # Create a ggplot scatter plot for breeding areas (Longitude_Start vs Latitude_Start)
  p <- ggplot() +
    geom_point(data = subset(plot_data, Population == "Chaun"),
               aes(x = Longitude_Start, y = Latitude_Start), color = "purple", shape = 16) +
    geom_point(data = subset(plot_data, Population == "Indigirka"),
               aes(x = Longitude_Start, y = Latitude_Start), color = "blue", shape = 17)

  # Add scatter points for wintering areas
  p <- p + geom_point(aes(x = stop_longitudes, y = stop_latitudes), color = "red", shape = 4)

  # Define the slope and intercept from the given equation
  slope <- 1.2919
  intercept <- -10.107

  # Generate x values from 0 to 180 with a step of 5
  x_values <- seq(0, 180, by = 5)

  # Calculate corresponding y values using the regression equation
  regression_y_values <- slope * x_values + intercept

  # Add the regression line to the plot
  p <- p + geom_line(aes(x = x_values, y = regression_y_values), color = "green")

  # Set axis limits and ticks
  p <- p + scale_x_continuous(limits = c(0, 180), breaks = seq(0, 180, by = 10)) +
    scale_y_continuous(limits = c(0, 90), breaks = seq(0, 90, by = 10))

  # Add labels and title
  p <- p + labs(x = "Longitude (Breeding Area)", y = "Longitude (Wintering Area)")

  # Add R² value and slope-intercept equation annotations
  r_squared <- 0.6624
  p <- p + annotate("text", x = 10, y = 85, label = sprintf("R² = %.4f", r_squared), size = 4) +
    annotate("text", x = 10, y = 80, label = sprintf("y = %.4fx + %.3f", slope, intercept), size = 4)

  # Show the plot
  print(p)

}, error = function(e) {
  cat("Error reading the Excel file:", conditionMessage(e), "\n")
})
