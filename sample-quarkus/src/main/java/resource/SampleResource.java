package resource;


import jakarta.annotation.Resource;
import jakarta.ws.rs.POST;
import jakarta.ws.rs.Path;
import jakarta.ws.rs.core.Response;
import lombok.extern.slf4j.Slf4j;

@Resource
@Path("/sample")
@Slf4j
public class SampleResource {

    @POST
    @Path("/test")
    public Response test(String message) {

        log.info("message: {}", message);
        return Response.ok().build();
    }
}
