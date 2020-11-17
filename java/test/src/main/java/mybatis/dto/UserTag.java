package mybatis.dto;

import lombok.Data;

import java.util.Date;

@Data
public class UserTag {
  private Integer id;
  private Integer userId;
  private String tag;
  private Date created;
  private Date updated;
}
