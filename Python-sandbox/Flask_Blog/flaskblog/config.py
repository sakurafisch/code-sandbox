import os
class Config:
   SECRET_KEY = '67f3196146d682809d7992a0351e2b26'
   SQLALCHEMY_DATABASE_URI = 'sqlite:///site.db'
   MAIL_SERVER = 'smtp.office365.com'
   MAIL_PORT = 587
   MAIL_USER_TLS = True
   MAIL_USERNAME = os.environ.get('EMAIL_USER')
   MAIL_PASSWORD = os.environ.get('EMAIL_PASS')