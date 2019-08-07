FROM python:3-alpine

# Update and install dependencies
RUN pip install --trusted-host pypi.org --trusted-host files.pythonhosted.org Flask BeautifulSoup4 flask-cors

# Add code
ADD .build/* /opt/webapp/

# Set the working directory
WORKDIR /opt/

# Set environment variables
ENV FLASK_APP=webapp
ENV FLASK_ENV=development

# Expose the application's port
EXPOSE 5000

# Run the application
CMD ["flask", "run", "--host=0.0.0.0"]
