package simple;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collection;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Map.Entry;
import java.util.Random;
import java.util.UUID;
import java.util.stream.Collectors;

import com.google.common.base.MoreObjects;
import com.google.common.collect.ArrayListMultimap;
import com.google.common.collect.HashMultimap;
import com.google.common.collect.ImmutableMultimap;
import com.google.common.collect.Lists;
import com.google.common.collect.Maps;
import com.google.common.collect.Multimap;
import com.google.common.collect.Multimaps;
import com.google.common.collect.SetMultimap;

import org.junit.Test;

public class StreamTest {

    @Test
    public void testFlatMap() {
        Map<String, List<Integer>> map = new HashMap<>();
        map.put("a", Lists.newArrayList(1, 2, 3));
        map.put("b", Lists.newArrayList(4, 5, 6));
        // want [1, 2, 3, 4, 5, 6]
        map.values().stream().flatMap(Collection::stream).forEach(System.out::println);
    }

    @Test
    public void testList2Map() {
        List<String> list = Lists.newArrayList("Aabc", "Bdef", "Cghi");
        // want {A: abc, B:def, C:ghi}
        Map<String, String> map = list.stream()
                .collect(Collectors.toMap(item -> item.substring(0, 1), item -> item.substring(1)));
        System.out.println(map);
    }

    @Test
    public void testList2MultiMap() {
        List<String> list = Lists.newArrayList("Aabc", "Adef", "Bghi");
        // want {A: [abc, def], B:[ghi]}
        Map<String, List<String>> map = list.stream()
                .collect(Collectors.toMap(item -> item.substring(0, 1), item -> Lists.newArrayList(item.substring(1)),
                        // key 冲突时，两个value的处理
                        (o, n) -> {
                            o.addAll(n);
                            return o;
                        }));

        System.out.println(map);
    }

    @Test
    public void testMultiMapTransform() {
        Map<String, List<String>> map = new HashMap<>();
        map.put("a", Lists.newArrayList("A", "B", "C"));
        // want {A: [a, b, c]}
        map = map.entrySet().stream().map(entry -> {
            String newKey = entry.getKey().toUpperCase();
            List<String> newVal = entry.getValue().stream().map(item -> item.toLowerCase())
                    .collect(Collectors.toList());
            // or new AbstractMap.SimpleEntry<String, String>(newKey, newVal);
            return Maps.immutableEntry(newKey, newVal);
        }).collect(Collectors.toMap(Entry::getKey, Entry::getValue));
        System.out.println(map);
    }

    // 使用 guava 可以很简洁的进行转换
    @Test
    public void testFlatMapGuava() {
        Multimap<String, Integer> multimap = ArrayListMultimap.create();
        multimap.putAll("a", Lists.newArrayList(1, 2, 3));
        multimap.putAll("b", Lists.newArrayList(4, 5, 6));
        // want [1, 2, 3, 4, 5, 6]
        System.out.println(multimap.values());
    }

    @Test
    public void testList2MapGuava() {
        List<String> list = Lists.newArrayList("Aabc", "Bdef", "Cghi");
        Map<String, String> map = Maps.uniqueIndex(list, item -> item.substring(0, 1));
        // want {A: abc, B:def, C:ghi}
        map = Maps.transformValues(map, item -> item.substring(1));
        System.out.println(map);
    }

    @Test
    public void testList2MultiMapGuava() {
        List<String> list = Lists.newArrayList("Aabc", "Adef", "Bghi");
        // want {A: [abc, def], B:[ghi]}
        Multimap<String, String> multimap = list.stream().collect(ArrayListMultimap::create,
                (m, item) -> m.put(item.substring(0, 1), item.substring(1)), Multimap::putAll);
        System.out.println(multimap);
    }

    @Test
    public void testMultiMapTransformGuava() {
        Multimap<String, String> multimap = ArrayListMultimap.create();
        multimap.putAll("a", Lists.newArrayList("A", "B", "C"));
        // want {A: [a, b, c]}
        // guava 不提供 transform key 的方法
        // https://stackoverflow.com/a/5733566/4276950
        multimap = Multimaps.transformValues(multimap, String::toLowerCase);
        System.out.println(multimap);
    }

    @Test
    public void testListGroupByTheEleOfNestedCollectionGuava() {
        // foo1, tags=a,b,c
        // foo2, tags=c,d
        // foo3, tags=a,c,e
        // want
        // a -> foo1, foo3
        // b -> foo1
        // c -> foo1, foo2, foo3
        // d -> foo2
        // e -> foo3
        List<String> list = Lists.newArrayList("foo1, tags=a,b,c", "foo2, tags=c,d", "foo3, tags=a,c,e");
        Multimap<String, String> multimap = ArrayListMultimap.create();
        list.forEach(item -> {
            String[] arr = item.split(", ");
            String[] tags = arr[1].substring("tags=".length()).split(",");
            Arrays.stream(tags).forEach(tag -> multimap.put(tag, arr[0]));
        });
        System.out.println(multimap);
    }

    @Test
    public void testListGroupByTheEleOfNestedCollectionGuava2() {
        List<String> list = Lists.newArrayList("foo1, tags=a,b,c", "foo2, tags=c,d", "foo3, tags=a,c,e");
        Multimap<String, String> map = list.stream()
                .collect(ImmutableMultimap.Builder<String, String>::new, (builder, item) -> {
                    String[] arr = item.split(", ");
                    String[] tags = arr[1].substring("tags=".length()).split(",");
                    Arrays.stream(tags).forEach(tag -> builder.put(tag, arr[0]));
                }, (builder1, builder2) -> builder1.putAll(builder2.build())).build();
        System.out.println(map);
    }

    @Test
    public void testMultiMapAsMap() {
        Multimap<String, A> cstr_aMap = cstr_aMap();
        // 将A中的blist和B中的clist只保留一个元素
        cstr_aMap = Multimaps.transformEntries(cstr_aMap, (cstr, a) -> {
            a.blist.stream()
                .filter(b -> b.getCList().contains(cstr))
                .findFirst().map(b -> {
                    C c = new C();
                    c.str = cstr;
                    b.clist = Lists.newArrayList(c);
                    return b;
                }).ifPresent(b -> a.blist = Lists.newArrayList(b));
            return a;
        });
        System.out.println(cstr_aMap);
    }

    public Multimap<String, A> cstr_aMap() {
        List<String> cset = new ArrayList<>();
        int bsize = 5;
        int csize = 10;
        for (int i = 0; i < csize; i++) {
            cset.add(UUID.randomUUID().toString().substring(0, 5));
        }
        Random r = new Random();
        List<B> blist = new ArrayList<>();
        for (int i = 0; i < bsize; i++) {
            B b = new B();
            b.id = i;
            b.aid = i / 2;
            int toIndex = r.nextInt(csize);
            while (toIndex == 0) {
                toIndex = r.nextInt(csize);
            }
            int fromIndex = r.nextInt(toIndex);
            List<String> clist = cset.subList(fromIndex, toIndex);
            while (clist.isEmpty()) {
                clist = cset.subList(fromIndex, toIndex);
            }
            b.clist = clist.stream().map(cstr -> {
                C c = new C();
                c.str = cstr;
                return c;
            }).collect(Collectors.toList());
            blist.add(b);
        }

        Map<Integer, List<B>> map = blist.stream().collect(Collectors.groupingBy(b -> b.aid));
        // cstr -> aid
        SetMultimap<String, Integer> cstr_aid = HashMultimap.create();
        blist.stream().forEach(b -> b.clist.forEach(c -> cstr_aid.put(c.str, b.aid)));

        List<A> alist = map.entrySet().stream().map(entry -> {
            A a = new A();
            a.id = entry.getKey();
            a.blist = entry.getValue();
            return a;
        }).collect(Collectors.toList());
        System.out.println(alist);
        Map<Integer, A> aMap = Maps.uniqueIndex(alist, a -> a.id);

        return Multimaps.transformValues(cstr_aid, aMap::get);
    }

    class A {
        int id;
        List<B> blist;

        @Override
        public String toString() {
            return MoreObjects.toStringHelper(this)
                .add("id", id)
                .add("blist", blist)
                .toString();
        }
    }

    class B {
        int id;
        int aid;
        List<C> clist;

        @Override
        public String toString() {
            return MoreObjects.toStringHelper(this)
                .add("id", id)
                .add("aid", aid)
                .add("clist", clist)
                .toString();
        }

        public List<String> getCList() {
            return clist.stream().map(c -> c.str).collect(Collectors.toList());
        }
    }

    class C {
        String str;

        @Override
        public String toString() {
            return str;
        }
    }

}