package mybatis.dto;

import lombok.Data;

import java.util.Date;
import java.util.List;

@Data
public class User {
  private Integer id;
  private String name;
  private Date created;
  private Date updated;
  private List<UserTag> tagList;
}