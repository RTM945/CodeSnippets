package mybatis.mapper;

import org.apache.ibatis.annotations.Param;
import mybatis.dto.User;

import java.util.List;

public interface UserMapper {
  List<User> getUserById(@Param("userId") Integer userId);
}
