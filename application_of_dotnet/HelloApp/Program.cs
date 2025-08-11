var builder = WebApplication.CreateBuilder(args);
var app = builder.Build();

app.MapGet("/", () => "Welcome!");
app.MapGet("/how-are-you", () => "I am good, how about you?");

app.Run("http://0.0.0.0:8085");
