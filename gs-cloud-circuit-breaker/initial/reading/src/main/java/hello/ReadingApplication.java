package hello;

import reactor.core.publisher.Mono;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.reactive.function.client.WebClient;
import org.springframework.beans.factory.annotation.Autowired;


@RestController
@SpringBootApplication
public class ReadingApplication {

    @Autowired
    private BookService bookService;

    @RequestMapping("/to-read")
    public Mono<String> toRead() {
      return WebClient.builder().build()
      .get().uri("http://localhost:8090/recommended").retrieve()
      .bodyToMono(String.class);
  }



  @RequestMapping("/to-read2")
  public Mono<String> toRead2() {
    return bookService.readingList();
  }

  public static void main(String[] args) {
    SpringApplication.run(ReadingApplication.class, args);
  }
}
