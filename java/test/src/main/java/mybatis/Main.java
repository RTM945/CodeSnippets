package mybatis;

import java.sql.ResultSet;
import java.util.List;

import org.apache.ibatis.mapping.MappedStatement;
import org.apache.ibatis.session.Configuration;
import org.apache.ibatis.session.RowBounds;
import org.jooq.DSLContext;
import org.jooq.Result;
import org.jooq.SQLDialect;
import org.jooq.impl.DSL;

public class Main {
    public static void main(String[] args) throws Exception {
        String json = "[{\"user_id\":5294147,\"user_name\":\"rtm\",\"user_created\":\"2020-11-11 12:27:41\",\"user_updated\":\"2020-11-11 12:27:41\",\"tag_id\":10002,\"tag_user_id\":5294147,\"tag\":\"lazy\",\"tag_created\":\"2020-11-11 12:27:41\",\"tag_updated\":\"2020-11-11 12:27:41\"},{\"user_id\":5294147,\"user_name\":\"rtm\",\"user_created\":\"2020-11-11 12:27:41\",\"user_updated\":\"2020-11-11 12:27:41\",\"tag_id\":10003,\"tag_user_id\":5294147,\"tag\":\"weak\",\"tag_created\":\"2020-11-11 12:27:41\",\"tag_updated\":\"2020-11-11 12:27:41\"}]";
        DSLContext dsl = DSL.using(SQLDialect.MYSQL);
        Result<?> result = dsl.fetchFromJSON(json);
        ResultSet rs = result.intoResultSet();

        // mybaits Configuration
        Configuration dummyConfiguration = new Configuration();
        dummyConfiguration.addMappers("mybatis.mapper");
        MappedStatement ms = dummyConfiguration.getMappedStatement("getUserById");

        FakeResultSetHandler resultHandler = new FakeResultSetHandler(ms, null, RowBounds.DEFAULT);
        List<Object> list = resultHandler.handleResultSets(rs);
        System.out.println(list);
    }
}
