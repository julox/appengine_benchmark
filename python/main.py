
import webapp2


class MainPage(webapp2.RequestHandler):
    def get(self):
        sum = 0
        for j in range(100):
          for i in range(1000000):
             sum += i
        self.response.headers['Content-Type'] = 'text/plain'
        self.response.write("Hello World!")


app = webapp2.WSGIApplication([
    ('/', MainPage),
], debug=True)
