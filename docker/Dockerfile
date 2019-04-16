# Comments in Dockerfiles
FROM python:3.7

# Update and install dependencies
RUN apt-get update
RUN pip install Flask

# Add code
ADD .build/* /opt/webapp/

# Set the working directory
WORKDIR /opt/webapp

# Set environment variables
ENV FLASK_APP=shop-backend
ENV FLASK_ENV=development

# Expose the application's port
EXPOSE 5000

# Run the application
CMD ["flask", "run", "--host=0.0.0.0"]